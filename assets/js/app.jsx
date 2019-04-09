
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
                <input type="text" value={this.state.firstname} onChange={this.handleFirstnameChange} />
                <input type="text" value={this.state.lastname} onChange={this.handleLastnameChange} />
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
        this.state = { sprint: props.sprint, showModal: false, modalDay: 0 };
        this.chooseMember = this.chooseMember.bind(this);
        this.saveSprint = this.saveSprint.bind(this);
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
        let s = this.props.sprint;
        s.days[this.state.modalDay] = memberId;
        this.setState({ sprint: s })
    };

    saveSprint() {
        axios
            .post("/sprints", this.state.sprint)
            .then((result) => {
                console.log('Sprint succesfully updated' + result);
            });
    }

    render() {
        const dateRow = this.props.sprint.days.map((day, i) => {
            let daysToAdd = i;
            if(i > 0) {
                daysToAdd += 2;
            }
            if(i > 5) {
                daysToAdd += 2;
            }
            let date = new Date(this.props.sprint.start);
            date.setDate(date.getDate() + daysToAdd);
            return (
                <th>{date.getDate() + ' ' + date.getMonth()}</th>
            );
        });

        const memberRow = this.props.sprint.days.map((day, i) => {
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
                    <tr><th colSpan={5}>Sprint nr: [{this.props.sprint.nr}] gestart op: [{this.props.sprint.start}]</th></tr>
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
                <button onClick={this.saveSprint}>Save sprint</button>
            </div>
        );
    }
}

class SprintCreator extends React.Component {
    constructor(props) {
        super(props);
        this.state = {nr: props.sprint.nr, start: props.sprint.start};

        this.handleNrChange = this.handleNrChange.bind(this);
        this.handleStartChange = this.handleStartChange.bind(this);
        this.createSprint = this.createSprint.bind(this);
    }

    handleNrChange(event) {
        console.log("Received new nr value: " + event.target.value);
        this.setState({nr: Number(event.target.value)});
    }

    handleStartChange(event) {
        this.setState({start: event.target.value});
    }

    componentWillReceiveProps(nextProps) {
        this.setState({nr: nextProps.sprint.nr, start: nextProps.sprint.start});
    }

    createSprint() {
        console.log(this.state);
        axios
            .put("/sprints", {nr: this.state.nr, start: this.state.start})
            .then((result) => {
                console.log('New sprint successfuly created' + result);
            })
            .catch((err) => {
                console.error(err);
            })
    }

    render() {
        return (
            <form onSubmit={this.createSprint}>
                <input type="number" value={this.state.nr} onChange={this.handleNrChange} />
                <input type="text" value={this.state.start} onChange={this.handleStartChange} />
                <input type="submit" value="Submit" />
            </form>
        )
    }
}

class Container extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            members: [],
            sprint: {days: []}
        }
    }

    componentDidMount() {
        axios
            .get("/members")
            .then((result) => {
                console.log('Received members: ' + result.data);
                this.setState({ members: result.data });
            });

        axios
            .get("/sprints")
            .then((result) => {
                console.log('Received sprint: ' + result.data);
                this.setState({ sprint: result.data });
            });
    }

    render() {
        return (
            <div>
                <MembersList members={this.state.members}/>
                <Sprint members={this.state.members} sprint={this.state.sprint}/>
                <SprintCreator sprint={this.state.sprint}/>
            </div>
        );
    }
}

ReactDOM.render( <Container/>, document.querySelector("#root"));
