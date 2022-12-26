import {Member} from "./Members";

export interface Rooms {
  roomName: string;
  members: Array<Member>;
  count: number;
  isVisible: boolean;
}
