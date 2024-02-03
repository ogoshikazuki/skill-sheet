<script setup lang="ts">
import { home } from '~/breadcrumbs'

definePageMeta({
  breadcrumbs: home()
})

const pages = [
  {
    path: 'basic-information',
    label: '基本情報'
  },
  {
    path: 'project',
    label: 'プロジェクト経歴'
  }
]

const externalPages = [
  {
    href: 'https://x.com/k_o_system',
    label: 'X'
  },
  {
    href: 'https://qiita.com/ogoshikazuki',
    label: 'Qiita'
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
      <v-list-item v-for="page in externalPages" :key="page.href" :href="page.href" target="_blank">
        {{ page.label }} <v-icon icon="mdi-open-in-new" />
      </v-list-item>
    </v-list>
  </v-card>
</template>
