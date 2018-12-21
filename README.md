# web-app created using gin, gorm, mysql

To implement a simple TODO web app. Here are the list of functions the app needs to support:

POST /v1/notes/create - create a note \n
GET /v1/notes - get the list of notes \n
GET /v1/notes/{note_id} - get a specific note
PATCH /v1/notes/{note_id} - Update a specific note
DEL /v1/notes/{note_id} - Delete a specific note
DEL /v1/notes - delete all notes

Every note has the following attributes: note_name, note_details, note_created_at, note_updated_at
