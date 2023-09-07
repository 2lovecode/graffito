<template>
    <el-row>
        <el-col :span="14">
            <div id="editor" style="width: 100%; height: 100vh;"></div>
        </el-col>
       
        <el-col :span="6" :offset="1">
            <el-button id="exec" @click="run" round>执行</el-button>
            <el-input
                type="textarea"
                placeholder="请输入内容"
                v-model="outputData"
                rows="10"
                show-word-limit
                >
            </el-input>
        </el-col>
    </el-row>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { sandboxRun } from '../api'
import * as monaco from 'monaco-editor/esm/vs/editor/editor.api'

let editor = null as monaco.editor.IStandaloneCodeEditor | null

export default defineComponent({
  data() {
    return {
      outputData: '',
      selectedLanguage: 'go' // 默认语言
    }
  },
  methods: {
    async run() {
      try {
        let code = ''
        if (editor) {
            code = editor?.getValue()
        }
        const data = await sandboxRun({
          sourceCode: code,
        });
        this.outputData = data.content;
      } catch (error) {
        console.error('Error running code:', error)
      }
    }
  },
  mounted() {
    // 创建编辑器实例
    
    const editorElement = document.getElementById("editor") as HTMLElement;

    editor = monaco.editor.create(editorElement, {
      value: 'println("hello world!")',
      language: this.selectedLanguage
    });

    // 添加其他编辑器配置和操作
    // ...

    // 监听编辑器事件
    // editor.onDidChangeModelContent(event => {
        
        
    // });
  },
  beforeUnmount() {
    editor?.dispose()
  }
});
</script>