import { toSimpleDateString } from "../../utils/date/converters.util"

export class CalendarEvent {
    id: string
    title: string
    description: string
    color: "red" | "yellow" | "blue"
    with?: string
    time: { start: string; end: string }
  
    constructor(
      title: string,
      description: string,
      timeStart: Date,
      timeEnd: Date,
      withPerson?: string,
      color: "red" | "yellow" | "blue" = "red"
    ) {
      this.title = title
      this.description = description
      this.time = {
        start: toSimpleDateString(timeStart),
        end: toSimpleDateString(timeEnd),
      }
      this.with = withPerson
      this.color = color
      this.id = new Date().getTime().toString()
    }
}