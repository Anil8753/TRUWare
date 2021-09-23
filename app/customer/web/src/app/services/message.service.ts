import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MessageService {
    private subject = new Subject<MessagePayload>();

    constructor() { }

    sendMessage(message: MessagePayload) {
        this.subject.next(message);
    }

    clearMessages() {
        this.subject.next();
    }

    getMessage(): Observable<MessagePayload> {
        return this.subject.asObservable();
    }
}

export enum MessageType{
  ClearFilter,
  LocationFilter,
  RatingsFilter,
  RateFilter,
  SpaceFilter,
}

export class MessagePayload {
  type: MessageType;
  data: any;
}