/*
 * CloudLinux Software Inc 2019-2021 All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { NO_ERRORS_SCHEMA } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormBuilder } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { ResetServiceUserPasswordDialogComponent } from '@ui/reset-service-user-password-dialog/reset-service-user-password-dialog.component';
import { MessagesService } from '@services/messages.service';
import { ServiceUsersService } from '@services/service-users.service';
import { ServicesPageService } from '@services/services-page.service';
import { MockMatDialogRef } from '@testing/mock-mat-dialog-ref';
import { MockMessageService } from '@testing/mock-messages-service';
import { MockServiceUsersService } from '@testing/mock-service-users-service';
import { MockServicesPageService } from '@testing/mock-services-page-service';

describe('ResetServiceUserPasswordDialogComponent', () => {
    let component: ResetServiceUserPasswordDialogComponent;
    let fixture: ComponentFixture<ResetServiceUserPasswordDialogComponent>;
    let serviceUsersService: MockServiceUsersService;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [ResetServiceUserPasswordDialogComponent],
            providers: [
                FormBuilder,
                { provide: MessagesService, useClass: MockMessageService },
                { provide: ServicesPageService, useClass: MockServicesPageService },
                { provide: ServiceUsersService, useClass: MockServiceUsersService },
                { provide: MatDialogRef, useClass: MockMatDialogRef },
                { provide: MAT_DIALOG_DATA, useValue: 'some_username' },
            ],
            schemas: [NO_ERRORS_SCHEMA]
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(ResetServiceUserPasswordDialogComponent);
        component = fixture.componentInstance;
        // @ts-ignore
        serviceUsersService = TestBed.inject(ServiceUsersService);
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });

    describe('when form is invalid', () => {
        beforeEach(() => {
            component.formGroup.patchValue({password: ''});
            fixture.detectChanges();
        });

        it('should not save user', () => {
            const spy = spyOn(serviceUsersService, 'editUser').and.callThrough();
            component.onSave();
            fixture.detectChanges();

            expect(spy).not.toHaveBeenCalled();
        });
    });

    describe('when form is valid', () => {
        beforeEach(() => {
            component.formGroup.patchValue({password: 'p@$$w0rd'});
            fixture.detectChanges();
        });

        it('should save user', () => {
            const spy = spyOn(serviceUsersService, 'editUser').and.callThrough();
            component.onSave();
            fixture.detectChanges();

            expect(spy).toHaveBeenCalled();
        });
    });
});
