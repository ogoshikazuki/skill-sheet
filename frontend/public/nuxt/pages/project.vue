<script setup lang="ts">
import { Temporal } from 'temporal-polyfill'
import { projects } from '~/breadcrumbs'
import { useProjectsQuery } from '~/graphql'

definePageMeta({
  breadcrumbs: projects()
})

const { result, loading } = useProjectsQuery()

const convertYearMonthForDisplay = (yearMonth: string) => {
  const plainYearMonth = Temporal.PlainYearMonth.from(yearMonth)
  return `${plainYearMonth.year}年${plainYearMonth.month}月`
}
</script>

<template>
  <div>
    <template v-if="loading">
      <v-card v-for="(project, i) in Array(3)" :key="i" :class="{ 'mt-2': i > 0 }">
        <v-card-item>
          <v-card-title>
            <v-skeleton-loader type="heading" />
          </v-card-title>
        </v-card-item>
        <v-card-text>
          <v-skeleton-loader type="text" />
        </v-card-text>
      </v-card>
    </template>
    <v-card v-for="(project, i) in result?.projects" v-else :key="project.id" :class="{ 'mt-2': i > 0 }">
      <v-card-item>
        <v-card-title :title="project.name" style="white-space: normal;">
          {{ project.name }}
        </v-card-title>
      </v-card-item>
      <v-card-text>
        <v-list>
          <v-list-item>
            <v-chip v-for="technology in project.technologies" :key="technology.id" class="ma-1">
              {{ technology.name }}
            </v-chip>
          </v-list-item>
          <v-list-item>
            {{ convertYearMonthForDisplay(project.startMonth) }} ~ {{ project.endMonth === null ? "" : convertYearMonthForDisplay(project.endMonth) }}
          </v-list-item>
        </v-list>
      </v-card-text>
    </v-card>
  </div>
</template>
