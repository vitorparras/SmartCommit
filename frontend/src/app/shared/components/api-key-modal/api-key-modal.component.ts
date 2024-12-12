import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-api-key-modal',
  templateUrl: './api-key-modal.component.html',
  styleUrls: ['./api-key-modal.component.scss']
})
export class ApiKeyModalComponent {
  apiKey: string = '';

  constructor(public dialogRef: MatDialogRef<ApiKeyModalComponent>) {}

  onSubmit() {
    this.dialogRef.close(this.apiKey);
  }

  onCancel() {
    this.dialogRef.close();
  }
}

