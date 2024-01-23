<script setup lang="ts">
const pages = [
  {
    path: 'basic-information',
    label: '基本情報'
  }
]
</script>

<template>
  <v-card>
    <v-list>
      <v-list-subheader>Menu</v-list-subheader>
      <v-list-item v-for="page in pages" :key="page.path" @click="navigateTo(page.path)">
        {{ page.label }}
      </v-list-item>
    </v-list>
  </v-card>
</template>
