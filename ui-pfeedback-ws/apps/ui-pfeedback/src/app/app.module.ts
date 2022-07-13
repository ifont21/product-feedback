import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AuthHttpInterceptor, AuthModule } from '@auth0/auth0-angular';
import { environment } from '../environments/environment';

import { AppComponent } from './app.component';
import { PFeedbackShellModule } from '@pfeedback/shell';
import { AppRoutingModule } from './app.routes';
import { BasePath } from '@pfeedback/board-data-spec';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    PFeedbackShellModule,
    AppRoutingModule,
    HttpClientModule,
    AuthModule.forRoot({
      ...environment.auth,
      httpInterceptor: {
        allowedList: [`${environment.dev.serverUrl}/api/*`],
      },
    }),
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthHttpInterceptor,
      multi: true,
    },
    {
      provide: BasePath,
      useValue: `${environment.dev.serverUrl}/api`,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
