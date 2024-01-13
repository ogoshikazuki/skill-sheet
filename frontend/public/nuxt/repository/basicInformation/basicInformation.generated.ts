import * as Types from '../types';

import gql from 'graphql-tag';
import * as VueApolloComposable from '@vue/apollo-composable';
import * as VueCompositionApi from '@vue/composition-api';
export type ReactiveFunction<TParam> = () => TParam;
export type BasicInformationQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type BasicInformationQuery = { __typename?: 'Query', basicInformation: { __typename?: 'BasicInformation', birthday: any } };


export const BasicInformationDocument = gql`
    query basicInformation {
  basicInformation {
    birthday
  }
}
    `;

/**
 * __useBasicInformationQuery__
 *
 * To run a query within a Vue component, call `useBasicInformationQuery` and pass it any options that fit your needs.
 * When your component renders, `useBasicInformationQuery` returns an object from Apollo Client that contains result, loading and error properties
 * you can use to render your UI.
 *
 * @param options that will be passed into the query, supported options are listed on: https://v4.apollo.vuejs.org/guide-composable/query.html#options;
 *
 * @example
 * const { result, loading, error } = useBasicInformationQuery();
 */
export function useBasicInformationQuery(options: VueApolloComposable.UseQueryOptions<BasicInformationQuery, BasicInformationQueryVariables> | VueCompositionApi.Ref<VueApolloComposable.UseQueryOptions<BasicInformationQuery, BasicInformationQueryVariables>> | ReactiveFunction<VueApolloComposable.UseQueryOptions<BasicInformationQuery, BasicInformationQueryVariables>> = {}) {
  return VueApolloComposable.useQuery<BasicInformationQuery, BasicInformationQueryVariables>(BasicInformationDocument, {}, options);
}
export function useBasicInformationLazyQuery(options: VueApolloComposable.UseQueryOptions<BasicInformationQuery, BasicInformationQueryVariables> | VueCompositionApi.Ref<VueApolloComposable.UseQueryOptions<BasicInformationQuery, BasicInformationQueryVariables>> | ReactiveFunction<VueApolloComposable.UseQueryOptions<BasicInformationQuery, BasicInformationQueryVariables>> = {}) {
  return VueApolloComposable.useLazyQuery<BasicInformationQuery, BasicInformationQueryVariables>(BasicInformationDocument, {}, options);
}
export type BasicInformationQueryCompositionFunctionResult = VueApolloComposable.UseQueryReturn<BasicInformationQuery, BasicInformationQueryVariables>;