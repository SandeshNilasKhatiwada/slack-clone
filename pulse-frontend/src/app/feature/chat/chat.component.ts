import { Component, inject, OnInit } from '@angular/core';
import { ChatService } from '../../core/services/chat.service';


@Component({
  selector: 'app-chat',
  standalone: true,
  templateUrl: './chat.component.html',
})
export class ChatComponent implements OnInit {
  private chatService = inject(ChatService);

  ngOnInit() {
    // Open the WebSocket connection the moment the user hits the chat page
    this.chatService.connect();

    // Send a test message 1 second after connecting
    setTimeout(() => {
      this.chatService.sendMessage("Hello from Angular!");
    }, 1000);
  }
}
