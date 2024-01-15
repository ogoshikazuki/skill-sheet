import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: '../../../graphql/*.graphqls',
  documents: './graphql/*.graphql',
  ignoreNoDocuments: true,
  generates: {
    './graphql/index.ts': {
      plugins: ['typescript', 'typescript-operations', 'typescript-vue-apollo']
    }
  }
}

export default config
