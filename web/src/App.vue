<script setup lang="ts">
import Aside from './components/layout/Aside.vue';
import Header from './components/layout/Header.vue';
import Footer from './components/layout/Footer.vue';

</script>
<template>
  <div>
    <el-container>
      <el-aside>
        <Aside />
      </el-aside>
      <el-container>
        <el-header>
          <Header />
        </el-header>
        <el-main>
          <router-view></router-view>
        </el-main>
        <el-footer>
          <Footer />
        </el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
</style>
