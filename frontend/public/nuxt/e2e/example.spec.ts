import { test, expect } from '@playwright/test'

test.beforeEach(async ({ page }) => {
  await page.goto('/')
})

test('basic information', async ({ page }) => {
  await page.getByText('基本情報').click()

  await expect(page.getByText('1991-07-01')).toBeVisible()
  await expect(page.getByText('男性')).toBeVisible()
  await expect(page.getByText('上智大学卒業')).toBeVisible()
})

test('project', async ({ page }) => {
  await page.getByText('プロジェクト経歴').click()

  await expect(page.getByText('健診PHR開発プロジェクト')).toBeVisible()
  await expect(page.getByText('2021年10月 ~')).toBeVisible()
  await expect(page.getByText('オンライン商談システムの管理画面保守開発')).toBeVisible()
  await expect(page.getByText('2020年7月 ~ 2021年3月')).toBeVisible()
  await expect(page.getByText('人材紹介会社向けクラウド型業務管理システムのリニューアル')).toBeVisible()
  await expect(page.getByText('2017年4月 ~ 2018年8月')).toBeVisible()
})
