import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: '../../../graphql/schema/*.graphqls',
  documents: './graphql/*.graphql',
  ignoreNoDocuments: true,
  generates: {
    './graphql/index.ts': {
      plugins: ['typescript', 'typescript-operations', 'typescript-vue-apollo'],
      config: {
        scalars: {
          Date: {
            input: 'string',
            output: 'string'
          }
        }
      }
    }
  }
}

export default config
