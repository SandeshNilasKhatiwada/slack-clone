import { Injectable } from '@angular/core';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

@Injectable({
  providedIn: 'root'
})
export class ChatService {

  private socket$!: WebSocketSubject<any>;

  public connect() {
    this.socket$ = webSocket('ws://localhost:8080/ws/chat');

    this.socket$.subscribe({
      next: (message) => console.log('Received message:', message),
      error: (err) => console.error('WebSocket error:', err),
      complete: () => console.log('WebSocket connection closed')
    })
  }
  public sendMessage(message: string) {
    if (this.socket$) {
      this.socket$.next({ content: message });
    } else {
      console.error('WebSocket connection is not established.');
    }
  }
}
