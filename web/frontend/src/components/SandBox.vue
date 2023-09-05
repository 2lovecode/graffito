<template>
    <el-row>
        <el-col :span="10">
            <el-input
                type="textarea"
                placeholder="请输入内容"
                v-model="inputData"
                rows="20"
                show-word-limit
                >
                </el-input>
        </el-col>
        <el-col :span="2">
            <el-button id="exec" @click="run" round>执行</el-button>
        </el-col>
        <el-col :span="8">
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

export default defineComponent({
    data() {
        return {
            inputData: '',
            outputData: ''
        }
    },
    methods: {
        run() {
            const me = this
            sandboxRun({
                sourceCode: me.inputData
            }).then(function(data){
                me.outputData = data.content
            })
        }
    }
})
</script>