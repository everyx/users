package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/adnaan/users/internal/models/workspaceinvitation"

	"github.com/adnaan/users/internal/models/userrole"

	"github.com/adnaan/users/internal/models/permission"

	"github.com/adnaan/users/internal/models/grouprole"

	"github.com/adnaan/users/internal/models/workspacerole"

	"github.com/google/uuid"

	"github.com/adnaan/users/internal/models"
)

type DefaultWorkspaceStore struct {
	Driver     string
	DataSource string
	Ctx        context.Context
	Client     *models.Client
}

func (d DefaultWorkspaceStore) GetWorkspace(id string) (string, string, map[string]interface{}, error) {
	oid, err := uuid.Parse(id)
	if err != nil {
		return "", "", nil, err
	}
	o, err := d.Client.Workspace.Get(d.Ctx, oid)
	if err != nil {
		return "", "", nil, err
	}

	return o.Name, o.Description, o.Metadata, nil
}

func (d DefaultWorkspaceStore) GetUserWorkspaces(userID string) (map[string][]string, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	roles, err := d.Client.WorkspaceRole.Query().Where(workspacerole.UserID(uid)).All(d.Ctx)
	if err != nil {
		return nil, err
	}
	rolemap := make(map[string][]string)
	for _, r := range roles {
		w, err := d.Client.Workspace.Get(d.Ctx, r.WorkspaceID)
		if err != nil {
			log.Printf("d.Client.Workspace.Get %v, err : %v", r.WorkspaceID, err)
			continue
		}
		data := []string{w.Name, r.Name}
		rolemap[r.WorkspaceID.String()] = data
	}
	return rolemap, nil
}

func (d DefaultWorkspaceStore) GetWorkspaceGroups(workspaceID string) ([]string, error) {
	wid, err := uuid.Parse(workspaceID)
	if err != nil {
		return nil, err
	}
	wrk, err := d.Client.Workspace.Get(d.Ctx, wid)
	if err != nil {
		return nil, err
	}
	wrkGroups, err := d.Client.Workspace.QueryGroups(wrk).All(d.Ctx)
	if err != nil {
		return nil, err
	}
	var grps []string
	for _, wg := range wrkGroups {
		grps = append(grps, wg.Name)
	}
	return grps, nil
}

func (d DefaultWorkspaceStore) CreateWorkspace(userID, name, description, plan string, metadata map[string]interface{}) (string, error) {
	var workspace *models.Workspace
	if err := withTx(d.Ctx, d.Client, func(tx *models.Tx) error {
		uid, err := uuid.Parse(userID)
		if err != nil {
			return err
		}
		workspace, err = tx.Workspace.Create().
			SetName(name).
			SetDescription(description).
			SetPlan(plan).
			SetMetadata(metadata).
			Save(d.Ctx)
		if err != nil {
			return err
		}

		_, err = tx.WorkspaceRole.Create().SetName("owner").SetWorkspaceID(workspace.ID).SetUserID(uid).Save(d.Ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}

	return workspace.ID.String(), nil
}

func (d DefaultWorkspaceStore) UpdateWorkspace(workspaceID, name, description, plan string, metadata map[string]interface{}) error {
	oid, err := uuid.Parse(workspaceID)
	if err != nil {
		return err
	}
	_, err = d.Client.Workspace.UpdateOneID(oid).
		SetName(name).
		SetPlan(plan).
		SetDescription(description).
		SetMetadata(metadata).
		Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d DefaultWorkspaceStore) DeleteWorkspace(workspaceID string) error {
	oid, err := uuid.Parse(workspaceID)
	if err != nil {
		return err
	}
	return d.Client.Workspace.DeleteOneID(oid).Exec(d.Ctx)
}

func (d DefaultWorkspaceStore) WorkspaceUpsertUser(workspaceID, userID, role string) error {
	oid, err := uuid.Parse(workspaceID)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	or, err := d.Client.WorkspaceRole.Query().Where(workspacerole.And(workspacerole.UserID(uid), workspacerole.WorkspaceID(oid))).First(d.Ctx)
	if err != nil {
		if !models.IsNotFound(err) {
			return err
		}
		_, err = d.Client.WorkspaceRole.Create().SetName(role).SetUserID(uid).SetWorkspaceID(oid).Save(d.Ctx)
		if err != nil {
			return err
		}
	}

	if or != nil {
		err = or.Update().SetName(role).Exec(d.Ctx)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (d DefaultWorkspaceStore) WorkspaceRemoveUser(workspaceID, userID string) error {
	oid, err := uuid.Parse(workspaceID)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	n, err := d.Client.WorkspaceRole.Delete().Where(workspacerole.And(workspacerole.UserID(uid), workspacerole.WorkspaceID(oid))).Exec(d.Ctx)
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("delete err : workspace %v role not found for user %v", oid, uid)
	}
	return nil
}

func (d DefaultWorkspaceStore) GetGroup(id string) (string, string, map[string]interface{}, error) {
	oid, err := uuid.Parse(id)
	if err != nil {
		return "", "", nil, err
	}
	t, err := d.Client.Group.Get(d.Ctx, oid)
	if err != nil {
		return "", "", nil, err
	}

	return t.Name, t.Description, t.Metadata, nil
}

func (d DefaultWorkspaceStore) GetUserGroupRoles(userID string) (map[string]string, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	roles, err := d.Client.GroupRole.Query().Where(grouprole.UserID(uid)).All(d.Ctx)
	if err != nil {
		return nil, err
	}
	rolemap := make(map[string]string)
	for _, r := range roles {
		rolemap[r.GroupID.String()] = r.Name
	}
	return rolemap, nil
}

func (d DefaultWorkspaceStore) CreateGroup(userID, workspaceID, name, description string, metadata map[string]interface{}) (string, error) {
	oid, err := uuid.Parse(workspaceID)
	if err != nil {
		return "", err
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return "", err
	}

	var tm *models.Group
	if err := withTx(d.Ctx, d.Client, func(tx *models.Tx) error {

		tm, err = tx.Group.Create().SetName(name).SetDescription(description).SetWorkspacesID(oid).SetMetadata(metadata).Save(d.Ctx)
		if err != nil {
			return err
		}

		_, err = tx.GroupRole.Create().SetName("maintainer").SetGroupID(tm.ID).SetUserID(uid).Save(d.Ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}

	return "", nil
}

func (d DefaultWorkspaceStore) UpdateGroup(groupID, name, description string, metadata map[string]interface{}) error {
	tid, err := uuid.Parse(groupID)
	if err != nil {
		return err
	}
	_, err = d.Client.Group.UpdateOneID(tid).SetName(name).SetDescription(description).SetMetadata(metadata).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d DefaultWorkspaceStore) DeleteGroup(groupID string) error {
	tid, err := uuid.Parse(groupID)
	if err != nil {
		return err
	}
	return d.Client.Group.DeleteOneID(tid).Exec(d.Ctx)
}

func (d DefaultWorkspaceStore) GroupUpsertUser(groupID, userID, role string) error {
	tid, err := uuid.Parse(groupID)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	tr, err := d.Client.GroupRole.Query().Where(grouprole.And(grouprole.UserID(uid), grouprole.GroupID(tid))).First(d.Ctx)
	if err != nil {
		if !models.IsNotFound(err) {
			return err
		}
		_, err = d.Client.GroupRole.Create().SetName(role).SetUserID(uid).SetGroupID(tid).Save(d.Ctx)
		if err != nil {
			return err
		}
	}

	if tr != nil {
		err = tr.Update().SetName(role).Exec(d.Ctx)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (d DefaultWorkspaceStore) GroupRemoveUser(groupID, userID string) error {
	tid, err := uuid.Parse(groupID)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	n, err := d.Client.GroupRole.Delete().Where(grouprole.And(grouprole.UserID(uid), grouprole.GroupID(tid))).Exec(d.Ctx)
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("delete err : group %v role not deleted for user %v", tid, uid)
	}
	return nil
}

func (d DefaultWorkspaceStore) GetUserRoles(userID string) ([]string, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	usrroles, err := d.Client.UserRole.Query().Where(userrole.UserID(uid)).All(d.Ctx)
	if err != nil {
		return nil, err
	}
	var roles []string
	for _, r := range usrroles {
		roles = append(roles, r.Name)
	}
	return roles, nil
}

func (d DefaultWorkspaceStore) CreateUserRole(userID, role string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	_, err = d.Client.UserRole.Create().SetName(role).SetUserID(uid).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d DefaultWorkspaceStore) DeleteUserRole(userID, role string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	n, err := d.Client.UserRole.Delete().
		Where(userrole.And(userrole.UserID(uid), userrole.Name(role))).Exec(d.Ctx)
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("delete err : userrole %v not deleted for userID %v", role, userID)
	}
	return nil
}

func (d DefaultWorkspaceStore) CreatePermission(roleID, action, target string) error {
	rid, err := uuid.Parse(roleID)
	if err != nil {
		return err
	}
	_, err = d.Client.Permission.Create().SetRoleID(rid).SetAction(action).SetTarget(target).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d DefaultWorkspaceStore) UpdatePermission(roleID, action, target string) error {
	rid, err := uuid.Parse(roleID)
	if err != nil {
		return err
	}
	_, err = d.Client.Permission.Update().
		Where(permission.And(permission.RoleID(rid), permission.Action(action))).SetTarget(target).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d DefaultWorkspaceStore) DeletePermission(roleID, action string) error {
	rid, err := uuid.Parse(roleID)
	if err != nil {
		return err
	}
	n, err := d.Client.Permission.Delete().
		Where(permission.And(permission.RoleID(rid), permission.Action(action))).Exec(d.Ctx)
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("delete err : permission action %v not deleted for roleID %v", action, roleID)
	}
	return nil
}

func (d DefaultWorkspaceStore) GetUserPermissions(userID string) (map[string]string, error) {
	allPermissions := make(map[string]string)
	// get all user roles
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	usrroles, err := d.Client.UserRole.Query().Where(userrole.UserID(uid)).All(d.Ctx)
	if err != nil {
		return nil, err
	}
	// permissions for all the user roles
	for _, usrrole := range usrroles {
		perms, err := d.Client.Permission.Query().Where(permission.RoleID(usrrole.ID)).All(d.Ctx)
		if err != nil {
			continue
		}
		for _, p := range perms {
			allPermissions[p.Action] = p.Target
		}
	}

	// get all group roles
	tmroles, err := d.Client.GroupRole.Query().Where(grouprole.UserID(uid)).All(d.Ctx)
	if err != nil {
		return nil, err
	}
	// permissions for all the group roles
	for _, tmrole := range tmroles {
		perms, err := d.Client.Permission.Query().Where(permission.RoleID(tmrole.ID)).All(d.Ctx)
		if err != nil {
			continue
		}
		for _, p := range perms {
			allPermissions[p.Action] = p.Target
		}
	}
	// get all workspace roles
	ogroles, err := d.Client.GroupRole.Query().Where(grouprole.UserID(uid)).All(d.Ctx)
	if err != nil {
		return nil, err
	}
	// permissions for all the workspace roles
	for _, ogrole := range ogroles {
		perms, err := d.Client.Permission.Query().Where(permission.RoleID(ogrole.ID)).All(d.Ctx)
		if err != nil {
			continue
		}
		for _, p := range perms {
			allPermissions[p.Action] = p.Target
		}
	}

	return allPermissions, nil
}

func (d DefaultWorkspaceStore) CreateInvitation(workspaceID, email, role string) error {
	wid, err := uuid.Parse(workspaceID)
	if err != nil {
		return err
	}
	_, err = d.Client.WorkspaceInvitation.Create().SetWorkspaceID(wid).SetEmail(email).SetRole(role).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d DefaultWorkspaceStore) DeleteInvitation(workspaceID, email string) error {
	wid, err := uuid.Parse(workspaceID)
	if err != nil {
		return err
	}
	d.Client.WorkspaceInvitation.Delete().Where(workspaceinvitation.And(
		workspaceinvitation.WorkspaceID(wid), workspaceinvitation.Email(email)))
	return nil
}

func (d DefaultWorkspaceStore) GetInvitations(email string) (map[string][]string, error) {
	wiArr, err := d.Client.WorkspaceInvitation.Query().Where(workspaceinvitation.Email(email)).All(d.Ctx)
	if err != nil {
		return nil, err
	}
	invitations := make(map[string][]string)
	for _, wi := range wiArr {
		ws, err := d.Client.Workspace.Get(d.Ctx, wi.ID)
		if err != nil {
			return nil, err
		}

		invitations[wi.WorkspaceID.String()] = []string{
			ws.Name,
			wi.Role,
		}
	}
	return invitations, nil
}
