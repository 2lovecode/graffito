<script setup lang="ts">
import { ref, onMounted } from 'vue'

import * as monaco from 'monaco-editor'
 
const editorContainer = ref<any>(null)
onMounted(() => {
  monaco.editor.create(editorContainer.value,{
    theme: "vs-dark",
    value: "",
    language:"go"
  })
})
</script>

<template>
  <div ref="editorContainer" class="editor"></div>
</template>

<style scoped>
.editor {
    width: 600px;
    flex: 1;
    min-height: 400px;
    overflow-y: auto;
  }
</style>