
// -*- JavaScript -*-


class MemberItem extends React.Component {
    constructor(props) {
        super(props);
        this.deleteMember = this.deleteMember.bind(this);
    }

    deleteMember() {
        axios
            .delete("/members", {
                params: {
                    id: this.props.id
                }
            })
    }

  render() {
    return (
      <tr>
        <td> {this.props.id}    </td>
        <td> {this.props.firstname} </td>
        <td> {this.props.lastname}  </td>
        <td> <button onClick={this.deleteMember}>Del</button> </td>
      </tr>
    );
  }
}

class MemberForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            firstname: '',
            lastname: ''
        };
        this.handleFirstnameChange = this.handleFirstnameChange.bind(this);
        this.handleLastnameChange = this.handleLastnameChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleFirstnameChange(event) {
        console.log('in firstname handler');
        this.setState({firstname: event.target.value});
    }

    handleLastnameChange(event) {
        this.setState({lastname: event.target.value});
    }

    handleSubmit(event) {
        console.info('A new member is added: ' + this.state.firstname + ' ' + this.state.lastname);
        event.preventDefault();
        axios
            .put("/members", null, {
                params: {
                    firstname: this.state.firstname,
                    lastname: this.state.lastname
                }
            })
            .then((result) => {
                console.log('Result of adding member: ' + result);
                this.setState({firstname: '', lastname: ''})
            });
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <input type="text" value={this.state.value} onChange={this.handleFirstnameChange} />
                <input type="text" value={this.state.value} onChange={this.handleLastnameChange} />
                <input type="submit" value="Submit" />
            </form>
        );
    }
}

class MembersList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const members = this.props.members.map((member, i) => {
      return (
        <MemberItem key={i} id={member.id} firstname={member.firstname} lastname={member.lastname} />
      );
    });

    return (
      <div>
        <table><tbody>
          <tr><th>Id</th><th>Firstname</th><th>Lastname</th><th>Delete</th></tr>
          {members}
        </tbody></table>
          <MemberForm/>

      </div>
    );
  }
}

class Modal extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        let members = this.props.members.map(m => {
            return (
                <tr onClick={this.props.handleChoice.bind(undefined, m.id)}>
                    <td> {m.id}    </td>
                    <td> {m.firstname} </td>
                    <td> {m.lastname}  </td>
                </tr>
            );
        });
        return (
            <div className={this.props.show ? "modal display-block" : "modal display-none"}>
                <section className="modal-main">
                    <div>
                        <table><tbody>
                        <tr><th>Id</th><th>Firstname</th><th>Lastname</th></tr>
                        {members}
                        </tbody></table>
                    </div>
                    <button onClick={this.props.handleClose}>close</button>
                </section>
            </div>
        );
    }
};

class Sprint extends React.Component {
    constructor(props) {
        super(props);
        this.state = { sprint: {Days: []}, showModal: false, modalDay: 0 };
        this.chooseMember = this.chooseMember.bind(this);
    }

    showModal = (day) => {
        console.log('Show modal: ' + day);
        this.setState({ showModal: true, modalDay: day });
    };

    hideModal = () => {
        this.setState({ showModal: false });
    };

    chooseMember(memberId) {
        console.log('Choose member: ' + memberId);
        console.log('Choose member for day: ' + this.state.modalDay);
        let s = this.state.sprint;
        s.Days[this.state.modalDay] = memberId;
        this.setState({ sprint: s })
    };

    componentDidMount() {
            axios
                .get("/sprints")
                .then((result) => {
                    console.log('Received sprint: ' + result.data);
                    var startDate = new Date(result.data.Start);
                    console.log('Set startdate: ' + startDate);
                    this.setState({ sprint: result.data, startDate: startDate });
                });
    }

    render() {
        const dateRow = this.state.sprint.Days.map((day, i) => {
            let daysToAdd = i;
            if(i > 0) {
                daysToAdd += 2;
            }
            if(i > 5) {
                daysToAdd += 2;
            }
            // console.log('adding amount: ' + daysToAdd);
            let date = new Date(this.state.startDate);
            date.setDate(date.getDate() + daysToAdd);
            return (
                <th>{date.getDate() + ' ' + date.getMonth()}</th>
            );
        });

        const memberRow = this.state.sprint.Days.map((day, i) => {
            let member = this.props.members.filter(m => m.id == day);
            return (
                <td onClick={this.showModal.bind(undefined, i)}>{member[0] ? member[0].firstname : ' '}</td>
            );
        });

        return (
            <div>
                <Modal show={this.state.showModal} handleClose={this.hideModal} handleChoice={this.chooseMember.bind(this)}
                       members={this.props.members}>
                </Modal>
                <table>
                    <thead>
                    <tr style={{width: '500px'}}>Sprint nr: [{this.state.sprint.Nr}] gestart op: [{this.state.sprint.Start}]</tr>
                    <tr>
                        <th>vrijdag</th><th>maandag</th><th>dinsdag</th><th>woensdag</th><th>donderdag</th><th>vrijdag</th><th>maandag</th><th>dinsdag</th><th>woensdag</th><th>donderdag</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        {dateRow}
                    </tr>
                    <tr>
                        {memberRow}
                    </tr>
                    </tbody>
                </table>
            </div>
        );
    }
}

class Container extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            members: []
        }
    }

    componentDidMount() {
        axios
            .get("/members")
            .then((result) => {
                console.log('Received members: ' + result.data);
                this.setState({ members: result.data });
            });
    }

    render() {
        return (
            <div>
                <MembersList members={this.state.members}/>
                <Sprint members={this.state.members}/>
            </div>
        );
    }
}

ReactDOM.render( <Container/>, document.querySelector("#root"));