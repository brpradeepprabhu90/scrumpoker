import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {ApiService} from "../api.service";
import {take} from "rxjs";
import {Messages} from "../models/Message";
import {WebSocketSubject} from 'rxjs/webSocket';
import {Rooms} from "../models/Rooms";
import {Member} from "../models/Members";
import {environment} from "../../environments/environment";

@Component({
  selector: 'app-poker',
  templateUrl: './poker.component.html',
  styleUrls: ['./poker.component.scss']
})
export class PokerComponent implements OnInit {
  roomName: string = "";
  userName: string = "";
  isShow = false;
  cards = [0, 1, 2, 3, 5, 8, 10, 12, 16, 20, 24, 32]
  points = 0;
  averagePoints = 0;
  usersList: Array<Member> = [];
  private socket$: WebSocketSubject<any> = new WebSocketSubject<any>(environment.socketPath);

  constructor(private router: Router, private activatedRoute: ActivatedRoute, private apiService: ApiService) {

  }

  ngOnInit() {
    const params = this.activatedRoute.snapshot.params
    this.roomName = params["roomId"]
    this.userName = params["userId"]
    console.log(this.roomName, this.userName)
    this.apiService.isUserPresent(this.roomName, this.userName).pipe(take(1)).subscribe({
      next: (data: Messages) => {
        this.getUsers();
      },
      error: (data: any) => {
        if (data.error?.Message) {
          if (data.error.Message === "Room " + this.roomName + " is not created") {
            this.router.navigate(["room"])
          } else {
            this.router.navigate([`room/${this.roomName}`])
          }
        }

      }
    })
    this.socket$.subscribe({
        next: message => {
          console.log(message);
          this.getUsers()
        },
        error: error => console.log('error:', error),
        complete: () => console.log('complete')
      }
    );
  }

  getUsers() {
    this.apiService.getUsers(this.roomName).pipe(take(1)).subscribe({
      next: (data: Rooms) => {
        if (this.isShow !== data.isVisible) {
          this.isShow = data.isVisible;
          if (!this.isShow) {
            this.points = 0;
          }
        }
        this.usersList = Object.values(data.members)
        const total = this.usersList.reduce((a: number, b: Member) => {
          return a + b.points
        }, 0)
        this.averagePoints = total / this.usersList.length
      },
      error: (data: any) => {
        console.log(data)

      }
    })
  }

  sendPoints(points: number) {
    this.points = points;
    this.socket$.next({
      message: {
        type: "updateUser",
        roomName: this.roomName,
        username: this.userName,
        points: points
      }
    });
  }

  onRevealCards() {
    this.socket$.next({
      message: {
        type: "revealCards",
        roomName: this.roomName,
        username: this.userName

      }
    });
  }

  onResetCards() {
    this.points = 0;
    this.socket$.next({
      message: {
        type: "resetCards",
        roomName: this.roomName,
        username: this.userName

      }
    });
  }
}
