<script lang="ts" setup>
import { Temporal } from 'temporal-polyfill'
import { useBasicInformationQuery } from '~/graphql'

const basicInformationQuery = useBasicInformationQuery()
const basicInformationQueryResult = basicInformationQuery.result
const basicInformations = computed(() => {
  const age = (() => {
    if (basicInformationQueryResult.value === undefined) {
      return ''
    }
    const birthday = Temporal.PlainDate.from(basicInformationQueryResult.value.basicInformation.birthday)
    const age = calculateAge(birthday, Temporal.Now.plainDateISO())
    return `(${age}歳)`
  })()
  return [
    {
      title: '生年月日',
      value: `${basicInformationQueryResult.value?.basicInformation.birthday}${age}`
    },
    {
      title: '性別',
      value: (basicInformationQueryResult.value === undefined) ? undefined : convertGenderForDisplay(basicInformationQueryResult.value.basicInformation.gender)
    },
    {
      title: '最終学歴',
      value: basicInformationQueryResult.value?.basicInformation.academicBackground
    }
  ]
})
</script>

<template>
  <v-card>
    <v-list>
      <v-list-subheader>
        基本情報
      </v-list-subheader>
      <v-skeleton-loader
        v-for="basicInformation in basicInformations"
        :key="basicInformation.title"
        :loading="basicInformationQuery.loading.value"
        type="list-item-two-line"
      >
        <v-list-item
          :title="basicInformation.title"
          :subtitle="basicInformation.value"
        />
      </v-skeleton-loader>
    </v-list>
  </v-card>
</template>
