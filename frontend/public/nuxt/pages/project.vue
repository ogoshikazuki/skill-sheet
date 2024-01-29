<script setup lang="ts">
import { Temporal } from 'temporal-polyfill'
import { useProjectsQuery } from '~/graphql'

const { result } = useProjectsQuery()

const convertYearMonthForDisplay = (yearMonth: string) => {
  const plainYearMonth = Temporal.PlainYearMonth.from(yearMonth)
  return `${plainYearMonth.year}年${plainYearMonth.month}月`
}
</script>

<template>
  <div>
    <v-card v-for="(project, i) in result?.projects" :key="project.id" :class="{ 'mt-2': i > 0 }">
      <v-card-item>
        <v-card-title :title="project.name" style="white-space: normal;">
          {{ project.name }}
        </v-card-title>
      </v-card-item>
      <v-card-text>
        {{ convertYearMonthForDisplay(project.startMonth) }} ~ {{ project.endMonth === null ? "" : convertYearMonthForDisplay(project.endMonth) }}
      </v-card-text>
    </v-card>
  </div>
</template>
