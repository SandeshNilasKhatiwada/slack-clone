import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

export interface AuthResponse {
  message: string;
  token?: string;
  user?: {
    id: number;
    username: string;
    email: string;
  }
}
@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private http = inject(HttpClient);
  private apiUrl = 'http://localhost:8080/api';

  register(userData: any) {
    return this.http.post<AuthResponse>(`${this.apiUrl}/user/register`, userData);
  }

  login(credentials: any) {
    return this.http.post<AuthResponse>(`${this.apiUrl}/user/login`, credentials);
  }

  // helper to save token to localStorage
  saveToken(token: string) {
    localStorage.setItem('authToken', token);
  }
}
