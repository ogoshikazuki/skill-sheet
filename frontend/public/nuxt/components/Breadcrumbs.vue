<script setup lang="ts">
const breadcrumbs = useBreadcrumbs()
</script>

<template>
  <v-breadcrumbs>
    <template v-for="(breadcrumb, i) in breadcrumbs" :key="breadcrumb.to">
      <v-breadcrumbs-divider v-if="i > 0" />
      <v-breadcrumbs-item :disabled="i + 1 === breadcrumbs.length" class="cursor-pointer" @click="navigateTo(breadcrumb.to)">
        {{ breadcrumb.title }}
      </v-breadcrumbs-item>
    </template>
  </v-breadcrumbs>
</template>
