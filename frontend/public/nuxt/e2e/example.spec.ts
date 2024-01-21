import { test, expect } from '@playwright/test'

test('basic information', async ({ page }) => {
  await page.goto('/')

  await page.getByRole('link', { name: '基本情報' }).click()

  await page.getByText('1991-07-01').waitFor()
  await expect(page.getByText('1991-07-01')).toBeVisible()
})
