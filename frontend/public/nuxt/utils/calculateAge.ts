import { Temporal } from 'temporal-polyfill'

const func = (birthday: Temporal.PlainDate, now: Temporal.PlainDate) => {
  return birthday
    .until(now)
    .round({ largestUnit: 'year', relativeTo: birthday })
    .years
}
export default func
