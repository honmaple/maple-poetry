<template>
  <q-layout view="lHh Lpr lFf">
    <q-header v-if="$q.screen.gt.xs">
      <q-toolbar>
        <router-link to="/">
          <q-avatar rounded>
            <q-img src="~assets/logo.png" :ratio="4/3" style="width: 100%;" />
          </q-avatar>
        </router-link>

        <q-space />
        <q-btn stretch flat label="诗集" :icon="icon.matAutoStories" to="/collections" />
        <q-btn stretch flat label="诗词" :icon="icon.matMenuBook" to="/poems" />
        <q-btn stretch flat label="诗人" :icon="icon.matPeople" to="/authors" />
        <q-btn stretch flat label="关于" :icon="icon.matContactSupport" to="/about" />
      </q-toolbar>
    </q-header>
    <q-footer v-else>
      <q-tabs no-caps indicator-color="transparent">
        <q-route-tab label="首页" :icon="icon.matHome" to="/" />
        <q-route-tab label="诗集" :icon="icon.matAutoStories" to="/collections" />
        <q-route-tab label="诗词" :icon="icon.matMenuBook" to="/poems" />
        <q-route-tab label="诗人" :icon="icon.matPeople" to="/authors" />
      </q-tabs>
    </q-footer>

    <q-page-container style="overflow-x: hidden;">
      <template v-if="$route.path == '/'">
        <q-page :style-fn="myStyle">
          <router-view :key="$route.path" />
        </q-page>
      </template>
      <template v-else>
        <q-page>
          <router-view :key="$route.path" />
        </q-page>
      </template>
    </q-page-container>
  </q-layout>
</template>

<script setup>
 import useIcon from '../composables/icon';

 const icon = useIcon()

 const myStyle = (offset, height) => {
     return {
         height: `calc(100vh - ${offset}px)`,
     }
 }
</script>
