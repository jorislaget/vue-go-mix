module.exports = {
    "prompts": {
        "name": {
            "type": "string",
            "required": true,
            "message": "Project name"
        },
        "description": {
            "type": "string",
            "required": false,
            "message": "Project description",
            "default": "A Vue.js + Go + Mix project"
        },
        "author": {
            "type": "string",
            "message": "Author"
        }
    },
    "completeMessage": "To get started:\n\n  {{^inPlace}}cd {{destDirName}}\n  cp .env.example .env\n  {{/inPlace}}npm install\n  npm run dev\n  go run main.go"
};