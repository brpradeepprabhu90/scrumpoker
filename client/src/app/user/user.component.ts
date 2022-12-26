import {Component, OnInit} from '@angular/core';
import {ApiService} from "../api.service";
import {MessageService} from "primeng/api";
import {ActivatedRoute, Router} from "@angular/router";
import {take} from "rxjs";
import {Messages} from "../models/Message";

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.scss']
})
export class UserComponent implements OnInit {
  userName: string = "";
  roomName: string = ""

  constructor(private apiService: ApiService, private messageService: MessageService, private router: Router, private activatedRoute: ActivatedRoute) {
  }

  ngOnInit() {
    this.roomName = this.activatedRoute.snapshot.params["roomId"]
    console.log(this.activatedRoute.snapshot.params)

  }

  onSubmit() {
    this.apiService.createUser(this.roomName, this.userName).pipe((take(1))).subscribe({
      next: (data: Messages) => {

        this.messageService.add({
          severity: "success",
          detail: `${this.userName} is created in ${this.roomName}`,
          summary: "User Creation"
        })
        setTimeout(() => {
          this.router.navigate([`room/${this.roomName}/${data.Message}`])
        }, 500)
      },
      error: (data: any) => {
        console.error("error=>", data.error)
        this.messageService.add({
          severity: "error",
          detail: data.error.Message,
          summary: "User Creation"
        })
      }
    })

  }
}
