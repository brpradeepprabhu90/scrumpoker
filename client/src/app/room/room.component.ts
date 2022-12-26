import {Component} from '@angular/core';
import {ApiService} from "../api.service";
import {take} from "rxjs";
import {MessageService} from "primeng/api";
import {Messages} from "../models/Message";
import {Router} from "@angular/router";

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent {
  roomName: string = "";

  constructor(private apiService: ApiService, private messageService: MessageService, private router: Router) {
  }

  onSubmit() {
    this.apiService.createRoom(this.roomName).pipe((take(1))).subscribe({
      next: (data: Messages) => {
        console.log(data)
        this.messageService.add({
          severity: "success",
          detail: data.Message,
          summary: "Room Creation"
        })
        this.router.navigate([`room/${this.roomName}`])
      },
      error: (data: any) => {
        console.error("error=>", data.error)
        this.messageService.add({
          severity: "error",
          detail: data.error.Message,
          summary: "Room Creation"
        })
      }
    })
  }
}
