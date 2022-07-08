import { CommonModule } from '@angular/common';
import { Component, NgModule, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-auth-button',
  template: `
    <h1>Product Feedback App</h1>
    <div>
      <h2>User Info</h2>
      {{ userProfile$ | async | json}}
    </div>
  `,
})
export class AuthButtonComponent implements OnInit {
  userProfile$!: Observable<unknown>;

  constructor(public authService: AuthService) {}

  ngOnInit(): void {
    this.userProfile$ = this.authService.user$;
  }
}

@NgModule({
  imports: [CommonModule],
  declarations: [AuthButtonComponent],
  exports: [AuthButtonComponent],
})
export class AuthButtonModule {}
