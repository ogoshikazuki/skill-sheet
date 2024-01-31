<template>
  <v-app>
    <v-app-bar color="grey-darken-3">
      <v-app-bar-title>
        KOシステム
      </v-app-bar-title>
      <v-btn icon="mdi-github" href="https://github.com/ogoshikazuki/skill-sheet" target="_blank" />
      <v-btn icon="mdi-home" @click="navigateToHome" />
    </v-app-bar>
    <v-main>
      <Breadcrumbs />
      <NuxtPage class="ma-4" />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
const navigateToHome = () => {
  navigateTo('/')
}
</script>
