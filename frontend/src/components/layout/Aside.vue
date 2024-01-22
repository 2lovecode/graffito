<script setup lang="ts">
import { ref } from 'vue'

const menu = ref([{'path': "/tools", 'name': "工具"}])
</script>

<template>
    <div v-for="(item, index) in menu" :key="index">
        <router-link :to="item.path">
            <span>{{ item.name }}</span>
        </router-link>
    </div>
</template>