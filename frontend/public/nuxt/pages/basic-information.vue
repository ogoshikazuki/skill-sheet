<script lang="ts" setup>
import { useBasicInformationQuery } from '~/graphql'

const basicInformationQuery = useBasicInformationQuery()
const basicInformationQueryResult = basicInformationQuery.result.value
const basicInformations = computed(() => {
  return [
    {
      title: '生年月日',
      value: basicInformationQueryResult?.basicInformation.birthday
    },
    {
      title: '性別',
      value: (basicInformationQueryResult === undefined) ? undefined : convertGenderForDisplay(basicInformationQueryResult.basicInformation.gender)
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
          v-skeleton-loader
          :title="basicInformation.title"
          :subtitle="basicInformation.value"
        />
      </v-skeleton-loader>
    </v-list>
  </v-card>
</template>
