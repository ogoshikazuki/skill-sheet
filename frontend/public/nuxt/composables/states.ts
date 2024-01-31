import type { Breadcrumbs } from '~/breadcrumbs/types'

export const useBreadcrumbs = () => useState<Breadcrumbs>('breadcrumbs', () => [])
