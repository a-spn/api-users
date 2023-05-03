package authorization_model

import "fmt"

type AuthorizationContext struct {
	SubjectID   uint
	SubjectRole string
	ObjectID    uint
	ObjectRole  string
	Method      string
}

func (auth *AuthorizationContext) String() string {
	return fmt.Sprintf("subject_id: %d, subject_role: %s, object_id: %d, object_role: %s,method: %s", auth.SubjectID, auth.SubjectRole, auth.ObjectID, auth.ObjectRole, auth.Method)
}
