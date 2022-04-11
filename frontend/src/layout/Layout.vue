<template>
  <el-container class="app-main">
    <div
      class="app-sidebar"
      data-wails-drag
    >
      <div class="app-ctrl" />
      <el-aside>
        <router-link
          v-for="(item, index) in menu"
          :key="index"
          :to="item.path"
        >
          <div :class="{ 'menu-item': true, 'menu-active': isActive(item.path) }">
            <el-icon>
              <render :content="item.icon" />
            </el-icon>
          </div>
        </router-link>
      </el-aside>
    </div>

    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import menu from '@/menu'
import { useRoute } from 'vue-router'
import { Render } from '@/components'

const route = useRoute()
const isActive = (path: string) => {
  return route.path === path
}
</script>

<style lang="scss">
$app-background-color: #e9eef3;
$app-sidebar-color: #d3dce6;
$app-sidebar-with: 68px;

html,
body,
#app {
  height: 100%;
}

.app-main {
  padding: 0px;
  margin: 0px;
  height: 100%;
  .app-title {
    background-color: $app-background-color;
    .app-title-ctrl {
      width: $app-sidebar-with;
      background-color: $app-sidebar-color;
    }
  }

  .app-sidebar {
    .app-ctrl {
      height: 28px;
    }
    background-color: $app-sidebar-color;
    .el-aside {
      width: 68px;
    }
  }

  .el-aside {
    background-color: $app-sidebar-color;
    color: var(--el-text-color-primary);
    text-align: center;
    line-height: 60px;
  }

  .el-main {
    background-color: $app-background-color;
    color: var(--el-text-color-primary);
  }
}

.menu-active {
  background-color: lightgray;
}
</style>
