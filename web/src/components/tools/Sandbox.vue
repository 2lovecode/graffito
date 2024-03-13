<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { sandboxRun } from '../../api'

import * as monaco from 'monaco-editor'

const output = ref('')

const editorContainer = ref<any>(null)

let editor = null as monaco.editor.IStandaloneCodeEditor | null

onMounted(() => {
  editor = monaco.editor.create(editorContainer.value,{
    theme: "vs-dark",
    value: "",
    language:"go"
  })
})
function run() {
  let code = editor?.getValue() == undefined ? "" : editor.getValue()
  sandboxRun({
    sourceCode: code,
  }).then((res) => {
    console.log(res)
    if (res != null) {
      if (res.data != null) {
        output.value = res.data.content == "" ? "无输出" : res.data.content
      } else {
        output.value = res.errorMessage
      }
    } else {
      output.value = "服务器错误"
    }
  }).catch((err) => {
    console.log(err)
  })
}
</script>

<template>
  <el-row>
    <el-col :span="12">
      <div ref="editorContainer" class="editor"></div>
    </el-col>
    <el-col :span="2" :offset="1">
      <el-button type="info" plain @click="run">Run</el-button>
    </el-col>
    <el-col :span="8">
      <el-input
        v-model="output"
        placeholder="输出结果"
        type="textarea"
        resize="none"
        rows="15"
      />
    </el-col>
  </el-row>
</template>

<style scoped>
.editor {
    width: 100%;
    flex: 1;
    min-height: 400px;
    overflow-y: auto;
  }
</style>