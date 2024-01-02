<template>
  <div>
    <h1>基本情報</h1>
    <ul>
      <li>年齢: {{ data?.basicInformation?.age }}</li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
// eslint-disable-next-line no-undef
const query = gql`
  query {
    basicInformation {
      age
    }
  }
`;

type BasicInformationResult = {
  basicInformation: {
    age: number;
  };
};

// eslint-disable-next-line no-undef
const { data } = await useAsyncQuery<BasicInformationResult>(query);
</script>
