<script setup lang="ts">
import { ref } from 'vue'
import sandboxSvg from '../../assets/tools/sandbox.svg'



const tools = ref([
    {
        name: "在线运行环境",
        path: "/tools/sandbox",
        icon: sandboxSvg,
        description: "在线运行环境，可以编写、运行、分享你的代码片段",
    }
])

</script>

<template>
    <el-row>
        <el-col v-for="(tool, index) in tools" :key="index" :span="6" :offset="index > 0 ? 2 : 0">
            <el-card :body-style="{ padding: '0px' }" @click="$router.push(tool.path)">
                <img :src="tool.icon">
                <div style="padding: 14px;text-align: center">
                    <span>{{ tool.name }}</span>
                </div>
            </el-card>
        </el-col>
    </el-row>
</template>