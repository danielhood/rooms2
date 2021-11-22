import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
    providedIn: 'root',
  })
export class MessageService {

    private messageSource = new BehaviorSubject(['Welcome to Rooms II', 'Your adventure awaits!', 'Please login to begin.']);
    currentMessage = this.messageSource.asObservable();

    constructor() { }

    setMessage(message: string) {
        this.messageSource.next([message])
    }

    setMessages(messages: string[]) {
        this.messageSource.next(messages)
    }

}
