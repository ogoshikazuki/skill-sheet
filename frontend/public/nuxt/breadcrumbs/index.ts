export const home = () => {
  return [{
    title: 'Home',
    to: '/'
  }]
}

export const basicInformation = () => {
  return [...home(), {
    title: '基本情報',
    to: '/basic-information'
  }]
}

export const projects = () => {
  return [...home(), {
    title: 'プロジェクト経歴',
    to: '/project'
  }]
}
