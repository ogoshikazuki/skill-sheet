import gql from 'graphql-tag';
import * as VueApolloComposable from '@vue/apollo-composable';
import * as VueCompositionApi from '@vue/composition-api';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type ReactiveFunction<TParam> = () => TParam;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Date: { input: any; output: any; }
};

export type BasicInformation = {
  __typename?: 'BasicInformation';
  birthday: Scalars['Date']['output'];
  gender: Gender;
};

export enum Gender {
  Female = 'FEMALE',
  Male = 'MALE'
}

export type Query = {
  __typename?: 'Query';
  basicInformation: BasicInformation;
};

export type BasicInformationQueryVariables = Exact<{ [key: string]: never; }>;


export type BasicInformationQuery = { __typename?: 'Query', basicInformation: { __typename?: 'BasicInformation', birthday: any, gender: Gender } };


export const BasicInformationDocument = gql`
    query basicInformation {
  basicInformation {
    birthday
    gender
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