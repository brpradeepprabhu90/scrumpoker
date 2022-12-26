import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {RoomComponent} from "./room/room.component";
import {UserComponent} from "./user/user.component";
import {PokerComponent} from "./poker/poker.component";

const routes: Routes = [
  {
    path: "room",
    component: RoomComponent

  },

  {
    path: "room/:roomId",
    component: UserComponent
  },

  {
    path: "room/:roomId/:userId",
    component: PokerComponent
  },
  {
    path: "",
    redirectTo: "room",
    pathMatch: 'full'

  },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
