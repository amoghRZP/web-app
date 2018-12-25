# web-app using gin, gorm, mysql

To implement a simple TODO web app. Here are the list of functions the app needs to support:

- POST /v1/notes/create - create a note 
- GET /v1/notes - get the list of notes 
- GET /v1/notes/{note_id} - get a specific note
- PATCH /v1/notes/{note_id} - Update a specific note
- DEL /v1/notes/{note_id} - Delete a specific note
- DEL /v1/notes - delete all notes

Every note has the following attributes: note_name, note_details, note_created_at, note_updated_at.


For each of the above routes, A simple basic auth via the header. To validate the auth we are using a custom middleware.
