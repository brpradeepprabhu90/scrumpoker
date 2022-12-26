import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {RoomComponent} from './room/room.component';
import {UserComponent} from './user/user.component';
import {PokerComponent} from './poker/poker.component';
import {InputTextModule} from 'primeng/inputtext';
import {CardModule} from "primeng/card";
import {ApiService} from "./api.service";
import {HttpClientModule} from "@angular/common/http";
import {MessageService} from "primeng/api";

import {FormsModule} from "@angular/forms";
import {CommonModule} from "@angular/common";
import {ToastModule} from "primeng/toast";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";

@NgModule({
  declarations: [
    AppComponent,
    RoomComponent,
    UserComponent,
    PokerComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule, BrowserAnimationsModule,
    InputTextModule,
    CardModule,
    HttpClientModule,
    ToastModule,
    FormsModule,

    CommonModule
  ],
  providers: [ApiService, MessageService],
  bootstrap: [AppComponent]
})
export class AppModule {
}
