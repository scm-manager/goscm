## TODO

JWT Token Auth

JWT User Login Auth?

Erstellen von SCM-User (Mittels Login / Einmal anmelden) (done)
--> in user.go

Hinzufügen in SCM-Gruppen

Löschen von SCM-User

Aus Gruppen entfernen

## Notizen

OpenAPI Spec: https://stagex.cloudogu.com/scm/openapi

Damit Tests ausgeführt werden können, muss ein gültiger Api Key in der
Umgebungsvariable `SCM_BEARER_TOKEN` verfügbar sein.

Tests werden ausgeführt mit SCM-Host: https://stagex.cloudogu.com/scm

## Features

### Group

- GetGroups() --> Liste aller Gruppen im SCMM
- GetGroup(groupID string) --> Gruppe mit der jeweiligen ID
- DeleteUserFromGroup(id userId, groupID string) --> loescht Nutzer aus angegebenen Gruppe
- AddUserToGroup(id userId, groupID string) --> Fügt Nutzer in angegebene Gruppe
- DeleteUserFromAllGroups(id userId) --> Entfernt Nutzer aus allen Gruppen im SCMM
- CopyGroupMembershipFromOtherUser(id userId, copyId userId)  --> Fügt id in alle Gruppen von CopyId

### User

- CreateUser(baseUrl string, username string, password string) --> erstellt den SCM-User
- DeleteUser(id string) löscht den übergebenen Nutzer
- DeleteUserAndGroupMembership(id userId) löscht den Nutzer und seine Memberships in den SCMM-Gruppen
