import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from "../environments/environment";
import {Observable} from "rxjs";
import {Messages} from "./models/Message";
import {Rooms} from "./models/Rooms";

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) {

  }

  createRoom(roomName: string): Observable<Messages> {
    return this.http.post<Messages>(`${environment.apiUrl}createRoom`, {roomName: roomName})
  }

  createUser(roomName: string, userName: string): Observable<Messages> {
    return this.http.post<Messages>(`${environment.apiUrl}createUser`, {
      roomName: roomName,
      userName: userName
    })
  }


  isUserPresent(roomName: string, userName: string): Observable<Messages> {
    return this.http.get<Messages>(`${environment.apiUrl}isUserPresent/${roomName}/${userName}`)
  }

  getUsers(roomName: string): Observable<Rooms> {
    return this.http.get<Rooms>(`${environment.apiUrl}getUsers/${roomName}`)
  }


}

