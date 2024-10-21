<template>
    <div id="app">
        <header>
            <h5>桌面便签</h5>
            <button @click="organizeNotes">整理</button>
            <!-- 监听 keyup 事件，按下回车键触发 addNote -->
            <input v-model="newNote" placeholder="添加新事项" @keyup.enter="addNote" />
            <!-- <div data-wails-drag> 拖动 </div> 拖不动-->
            <button @click="addNote">+</button>
        </header>

        <ul>
            <li v-for="(note, index) in notes" :key="index" class="note-item">
                <!-- 如果当前是编辑模式，显示输入框，否则显示编号 @blur="finishEdit(index, note)"-->
                <span :class="{ 'Done': note.Done }">
                    <template v-if="editIndex === index">
                        <input type="text" v-model.number="note.Rank"  @keyup.enter="finishEdit(index, note)"
                            style="width: 2ch; background: transparent; border: none; color: white;" />
                        <span @click="editIndex = index"> {{ note.Content }} </span>
                    </template>
                    <template v-else>
                        <span @click="editIndex = index">
                            {{ note.Rank }}. {{ note.Content }}
                        </span>
                    </template>
                </span>
                <div class="actions">
                    <input type="checkbox" v-model="note.Done" @click="done(index, note)" />
                    <span class="drag-handle">☰</span>
                    <button @click="deleteNote(index, note)">🗑️</button>
                </div>
            </li>
        </ul>
    </div>
</template>
  
<style scoped>
header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
}

.note-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin: 0.5rem 0;
}

.actions {
    display: flex;
    align-items: center;
}

.Done {
    text-decoration: line-through;
    /* 划线效果 */
    font-style: italic;
    /* 改为斜体 */
    color: gray;
    /* 可以设置颜色为灰色 */
}

.drag-handle {
    margin: 0 10px;
    cursor: grab;
}
</style>

<script>
import { GetNotes, AddNote, DeleteNote, Done,EditRank } from '../../wailsjs/go/dao/NoteDao';

// 用来退出程序
const handleQuit = () => {
    window.runtime.Quit();
}
export default {
    data() {
        return {
            editIndex: -1, // 存储当前正在编辑的事项编号的索引
            notes: [],  // 便签列表
            newNote: '',  // 新便签输入
        };
    },

    mounted: function () {
        this.loadNotes()
    },

    methods: {
        // 加载所有事项
        loadNotes() {
            GetNotes().then((result) => {
                this.notes = result;
            });
        },
        // 添加新事项
        addNote() {
            console.log(this.notes);
            if (this.newNote.trim()) {
                AddNote(this.newNote).then((result) => {
                    this.loadNotes();  // 等待 AddNote 完成后再加载事项
                    this.newNote = '';  // 清空输入框
                }).catch((error) => {
                    console.error('添加事项失败:', error);  // 处理错误
                });
                // 外面这里是异步的 then会先返回 promise对象
            }
        },
        // 完成事项
        done(index, note) {
            Done(note.Content).then((result) => {

            }).catch((error) => {

            })
        },
        // 删除事项
        deleteNote(index, note) {
            DeleteNote(note.Content).then((result) => {
                console.log(note);
                this.loadNotes();
            }).catch((error => {

            }))
        },
        organizeNotes() {
            this.loadNotes();
            console.log("Organizing notes..."); // 模拟整理功能
        },
        // 触发编辑完成的回调函数
        finishEdit(index, note) {
            console.log("==============");
            // 退出编辑模式
            this.editIndex = -1;
            EditRank(note).then((response)=>{
                this.loadNotes()
            }).catch((error)=>{

            })
        },
    },
};
</script>
  

  