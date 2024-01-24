import { Gender } from '../graphql'

const func = (gender: Gender) => {
  switch (gender) {
    case Gender.Male:
      return '男性'
    case Gender.Female:
      return '女性'
  }
}
export default func
