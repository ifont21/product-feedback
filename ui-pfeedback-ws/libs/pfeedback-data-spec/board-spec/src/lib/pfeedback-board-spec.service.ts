import { BoardDetailDTO } from './board-types';
import { Inject, Injectable, Optional } from '@angular/core';
import { SpecConfig, Configuration } from './config';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class PFeedbackBoardSpecService {
  constructor(
    @Optional() @Inject(SpecConfig) private config: Configuration,
    private http: HttpClient
  ) {}

  getBoardDetails(boardId: string): Observable<BoardDetailDTO> {
    return this.http.get<BoardDetailDTO>(
      `${this.config?.basePath}/boards/${boardId}`
    );
  }
}
