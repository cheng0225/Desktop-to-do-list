<template>
    <div id="app">
        <h1>Schedule Notes</h1>

        <!-- Add new note -->
        <div>
            <input v-model="newNote" placeholder="Add a new note" />
            <button @click="addNote">Add Note</button>
        </div>

        <!-- Display notes -->
        <ul>
            <li v-for="(note, index) in notes" :key="index">
                {{ note }}
                <button @click="deleteNote(index)">Delete</button>
            </li>
        </ul>
    </div>
</template>
  
<script>
export default {
    data() {
        return {
            newNote: '',  // to store the current input value
            notes: []     // to store the list of notes
        };
    },
    methods: {
        // Add a new note to the list
        addNote() {
            if (this.newNote.trim()) {
                this.notes.push(this.newNote);
                this.newNote = '';  // clear the input field
            }
        },
        // Delete a note from the list
        deleteNote(index) {
            this.notes.splice(index, 1);
        }
    }
};
</script>
  