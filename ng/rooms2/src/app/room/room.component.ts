import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { MessageService } from '../services/message.service'

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent implements OnInit, OnDestroy {
  messageSubscription!: Subscription;
  messages: string[] = [""];

  constructor(
    private messageService: MessageService
    ) {
    }

    ngOnInit() {
      this.messageSubscription = this.messageService.currentMessage.subscribe((messages: string[]) => this.messages = messages);
    }
  
    ngOnDestroy() {
      this.messageSubscription.unsubscribe();
    }
  
}
