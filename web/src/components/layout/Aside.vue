<script setup lang="ts">
import { ref } from 'vue'

const menu = ref([
    {
        'name': "首页",
        'index': "1",
        'icon': "House",
        'path': "/"
    },
    {
        'name': "菜单",
        'index': "2",
        'icon': "Menu",
        'path': "",
        'chirldren': [
            {
                'path': "/tools",
                'index': "2-1",
                'name': "工具",
            }
        ]
    }
    ])
</script>

<template>
    <el-row>
        <el-col>
            <h5>Graffito</h5>
            <el-menu
                default-active="1"
            >
                <dev v-for="(item, index) in menu" :key="index">
                    <el-menu-item v-if="item.chirldren == undefined" :index="item.index" @click="$router.push(item.path)">
                        <el-icon><component :is="item.icon" /></el-icon>
                        <template #title>{{ item.name }}</template>
                    </el-menu-item>
                    <el-sub-menu v-if="item.chirldren != undefined" :index="item.index">
                    <template #title>
                        <el-icon><component :is="item.icon" /></el-icon>
                        <span>{{ item.name }}</span>
                    </template>
                    <el-menu-item v-for="(child, index) in item.chirldren" :key="index" :index="child.index" @click="$router.push(child.path)">{{ child.name }}</el-menu-item>
                </el-sub-menu>
                </dev>
            </el-menu>
        </el-col>
    </el-row>
</template>