<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <router-link to="/">
          <q-avatar rounded>
            <q-img src="~assets/logo.png" :ratio="4/3" style="width: 100%;" />
          </q-avatar>
        </router-link>

        <q-space />
        <q-btn stretch flat label="诗集" icon="auto_stories" to="/collections" />
        <q-btn stretch flat label="诗词" icon="menu_book" to="/poems" />
        <q-btn stretch flat label="诗人" icon="people" to="/authors" />
        <q-btn stretch flat label="关于" icon="contact_support" to="/about" />
      </q-toolbar>
    </q-header>

    <!-- <q-footer class="bg-white text-black">
         <div class="flex flex-center">
         © 2023 honmaple
         </div>
         </q-footer>
    -->
    <q-page-container>
      <q-page style="max-width: 800px; margin: 0 auto;">
        <router-view :key="$route.path" />
      </q-page>
    </q-page-container>
  </q-layout>
</template>
