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
  Date: { input: string; output: string; }
  YearMonth: { input: any; output: any; }
};

export type BasicInformation = {
  __typename?: 'BasicInformation';
  academicBackground: Scalars['String']['output'];
  birthday: Scalars['Date']['output'];
  gender: Gender;
  id: Scalars['ID']['output'];
};

export enum Gender {
  Female = 'FEMALE',
  Male = 'MALE'
}

export enum OrderDirection {
  Asc = 'ASC',
  Desc = 'DESC'
}

export type Project = {
  __typename?: 'Project';
  endMonth?: Maybe<Scalars['YearMonth']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  startMonth: Scalars['YearMonth']['output'];
};

export type ProjectOrder = {
  direction: OrderDirection;
  field: ProjectOrderField;
};

export enum ProjectOrderField {
  EndMonth = 'END_MONTH',
  StartMonth = 'START_MONTH'
}

export type Query = {
  __typename?: 'Query';
  basicInformation: BasicInformation;
  projects: Array<Project>;
};


export type QueryProjectsArgs = {
  orderBy?: Array<ProjectOrder>;
};

export type BasicInformationQueryVariables = Exact<{ [key: string]: never; }>;


export type BasicInformationQuery = { __typename?: 'Query', basicInformation: { __typename?: 'BasicInformation', id: string, birthday: string, gender: Gender, academicBackground: string } };

export type ProjectsQueryVariables = Exact<{ [key: string]: never; }>;


export type ProjectsQuery = { __typename?: 'Query', projects: Array<{ __typename?: 'Project', id: string, name: string, startMonth: any, endMonth?: any | null }> };


export const BasicInformationDocument = gql`
    query BasicInformation {
  basicInformation {
    id
    birthday
    gender
    academicBackground
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
export const ProjectsDocument = gql`
    query Projects {
  projects {
    id
    name
    startMonth
    endMonth
  }
}
    `;

/**
 * __useProjectsQuery__
 *
 * To run a query within a Vue component, call `useProjectsQuery` and pass it any options that fit your needs.
 * When your component renders, `useProjectsQuery` returns an object from Apollo Client that contains result, loading and error properties
 * you can use to render your UI.
 *
 * @param options that will be passed into the query, supported options are listed on: https://v4.apollo.vuejs.org/guide-composable/query.html#options;
 *
 * @example
 * const { result, loading, error } = useProjectsQuery();
 */
export function useProjectsQuery(options: VueApolloComposable.UseQueryOptions<ProjectsQuery, ProjectsQueryVariables> | VueCompositionApi.Ref<VueApolloComposable.UseQueryOptions<ProjectsQuery, ProjectsQueryVariables>> | ReactiveFunction<VueApolloComposable.UseQueryOptions<ProjectsQuery, ProjectsQueryVariables>> = {}) {
  return VueApolloComposable.useQuery<ProjectsQuery, ProjectsQueryVariables>(ProjectsDocument, {}, options);
}
export function useProjectsLazyQuery(options: VueApolloComposable.UseQueryOptions<ProjectsQuery, ProjectsQueryVariables> | VueCompositionApi.Ref<VueApolloComposable.UseQueryOptions<ProjectsQuery, ProjectsQueryVariables>> | ReactiveFunction<VueApolloComposable.UseQueryOptions<ProjectsQuery, ProjectsQueryVariables>> = {}) {
  return VueApolloComposable.useLazyQuery<ProjectsQuery, ProjectsQueryVariables>(ProjectsDocument, {}, options);
}
export type ProjectsQueryCompositionFunctionResult = VueApolloComposable.UseQueryReturn<ProjectsQuery, ProjectsQueryVariables>;